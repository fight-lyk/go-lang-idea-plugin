/*
 * Copyright 2013-2016 Sergey Ignatov, Alexander Zolotov, Florin Patan
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.goide.editor.surround;

import com.goide.psi.GoFile;
import com.goide.psi.GoPsiTreeUtil;
import com.goide.psi.GoStatement;
import com.intellij.lang.surroundWith.SurroundDescriptor;
import com.intellij.lang.surroundWith.Surrounder;
import com.intellij.psi.PsiElement;
import com.intellij.psi.PsiFile;
import org.jetbrains.annotations.NotNull;

public class GoStatementsSurroundDescriptor implements SurroundDescriptor {
  private static final Surrounder[] SURROUNDERS = new Surrounder[]{
    new GoWithIfSurrounder(),
    new GoWithIfElseSurrounder(),
    new GoWithForSurrounder(),
    new GoWithBlockSurrounder()
  };

  @NotNull
  @Override
  public Surrounder[] getSurrounders() {
    return SURROUNDERS;
  }

  @NotNull
  @Override
  public PsiElement[] getElementsToSurround(PsiFile file, int startOffset, int endOffset) {
    return file instanceof GoFile
           ? GoPsiTreeUtil.getTopLevelElementsInRange((GoFile)file, startOffset, endOffset, GoStatement.class)
           : PsiElement.EMPTY_ARRAY;
  }

  @Override
  public boolean isExclusive() {
    return false;
  }
}
